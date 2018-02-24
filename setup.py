# -*- coding: utf-8 -*-
"""
 Created by wq.wang on 2017/8/25
"""
from setuptools import setup, find_packages

VERSION = open('VERSION').read().strip()

setup(
    name='ginDoc',
    version=VERSION,
    description='ginDoc web',
    author='wei.zhou',
    author_email='wei.zhou.wang@shunwang.com',
    packages=find_packages()
)
