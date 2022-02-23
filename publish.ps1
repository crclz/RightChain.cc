rm dist -Recurse -ErrorAction SilentlyContinue
python -m build
python -m twine upload dist/*