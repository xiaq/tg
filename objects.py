#!/usr/bin/python

import os
import sys
from collections import namedtuple

from lxml import html


type_map = {
    'Integer': 'int',
    'String': 'string',
    'True': 'bool',
    'Boolean': 'bool',
    'Float': 'float64',
}
all_caps = set(['id', 'url', 'mpeg4', 'gif'])


def gocase(s):
    return ''.join([w.upper() if w in all_caps else w.title()
                    for w in s.split('_')])


def parse_type(words, full, optional):
    if len(words) == 0:
        raise Exception('Empty type')
    if len(words) == 1:
        star = '*' if optional and words[0] != 'True' else ''
        return star + type_map.get(words[0], words[0])
    if words[0] == 'Float' and words[1] == 'number':
        words[1] = 'float64'
        return parse_type(words[1:], full, optional)
    if words[0] == 'Array' and words[1] == 'of':
        return '[]' + parse_type(words[2:], full, False)
    if words[1] == 'or':
        return type_map.get(words[0], words[0]) + '|' + parse_type(words[2:], full, False)
    else:
        raise Exception('Cannot parse type %s (full: %s)' % (words, full))


def main(out=sys.stdout, cache=''):
    if cache:
        with open(cache) as f:
            content = f.read()
    else:
        import requests
        content = requests.get('https://core.telegram.org/bots/api').content
    out.write('package tgbot\n\n'
              '// GENERATED AUTOMATICALLY BY objects.py\n\n')
    tree = html.document_fromstring(content)
    h4s = tree.xpath('//h4')
    for h4 in h4s:
        p = h4.xpath('following-sibling::*[1]')[0]
        s2 = h4.xpath('following-sibling::*[2]')[0]
        if p.tag != 'p' or s2.tag not in ('table', 'ul'):
            continue
        name = h4.xpath('text()')[0]
        if s2.tag == 'ul':
            # union definition
            ul = s2
            types = s2.xpath('li//text()')
            out.write('type %s interface{} // %s\n\n' % (name, '|'.join(types)))
        else:
            # method or object definition
            table = s2
            is_method = 'method' in ''.join(p.xpath('.//text()'))
            if is_method:
                name = name[0].upper() + name[1:] + 'Request'

            out.write('type %s struct {\n' % name)
            for tr in table.xpath('tbody/tr')[1:]:
                field_name = tr.xpath('td[1]/text()')[0]
                # This happens to work with both method tables and object tables
                # despite they having different formats, yay!
                optional = tr.xpath('td[3]//text()')[0].lower().startswith('optional')
                type_words = ' '.join(tr.xpath('td[2]//text()')).split()
                type_ = parse_type(list(type_words), type_words, optional)
                trailing = ''
                if '|' in type_:
                    trailing = ' // ' + type_
                    type_ = 'interface{}'
                out.write('\t%s %s `json:"%s%s"`%s\n' %
                          (gocase(field_name), type_, field_name,
                           ',omitempty' if optional and is_method else '', trailing))
            out.write('}\n\n')

if __name__ == '__main__':
    main(open('objects.go', 'w'), sys.argv[1] if len(sys.argv) >= 2 else '')
    os.system('gofmt -w objects.go')
