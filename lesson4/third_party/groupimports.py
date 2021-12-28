import re
import sys
from collections import defaultdict

assert len(sys.argv) == 2, "expected one argument"
file_name = sys.argv[1]
assert file_name.endswith(".go"), "expected go file"

with open(file_name) as f:
    content = [line for line in f]

first_import_index = content.index("import (\n") if "import (\n" in content else -1
if first_import_index == -1:
    exit(0)

imports = []
_in_imports = False
for line in content:
    if line == ")\n" and _in_imports:
        _in_imports = False

    if _in_imports:
        imports.append(line)

    if line == "import (\n":
        _in_imports = True

i = 0
while i < len(content) - 1:
    line = content[i]
    if line == ")\n" and _in_imports:
        content.pop(i)
        _in_imports = False

    if _in_imports:
        content.pop(i)
        continue

    if line == "import (\n":
        content.pop(i)
        _in_imports = True
        continue

    i += 1

imports = list(filter(lambda x: x != '\n', imports))

host_re = re.compile(r'\t.*"([a-zA-Z0-9-_]+\.)*[a-zA-Z0-9][a-zA-Z0-9-_]+\.[a-zA-Z]{2,11}?/.*')
common_imports = list(filter(lambda x: not host_re.match(x), imports))

uncommon_imports = list(filter(lambda x: host_re.match(x), imports))
uncommon_imports.sort(key=lambda x: re.sub(r'[^"]*"', '"', x, count=1))

if len(common_imports) != 0 and len(uncommon_imports) != 0:
    common_imports.append('\n')


def group_imports(imports, depth):
    if depth == 3:
        return imports

    if len(imports) < 6:
        return imports

    groups = defaultdict(list)
    for i in range(len(imports)):
        split = re.sub(r'[^"]*"', '"', imports[i], count=1).split('/')
        import_key = ""
        for j in range(len(split)):
            if j <= depth:
                import_key += split[j]

        groups[import_key].append(imports[i])

    res = []
    for group in groups.values():
        ggroup = group_imports(group, depth+1)

        if len(res) != 0:
            res.append('\n')

        for imp in ggroup:
            res.append(imp)

    return res


uncommon_imports = group_imports(uncommon_imports, 0)
imports = common_imports + uncommon_imports

content.insert(first_import_index, "import (\n")
content.insert(first_import_index + 1, ")\n")

imports.reverse()
for line in imports:
    content.insert(first_import_index + 1, line)


with open(file_name, 'w+') as f:
    f.writelines(content)
