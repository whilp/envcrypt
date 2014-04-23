from setuptools import setup

meta = dict(
    name="envcrypt",
    version="0.0.1",
    py_modules=["envcrypt"],
    author='Will Maier',
    author_email='will@simple.com',
    test_suite='tests',
    scripts=['scripts/envcrypt'],
    # entry_points=dict(
    #     console_scripts=[
    #         "envcrypt = envcrypt:run"
    #     ]
    # ),
)

setup(**meta)
