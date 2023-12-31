#!/bin/bash
declare -A LANGUAGE_ID_MAP=(
    ["python"]=1
    ["go"]=2
    ["typescript"]=3
    ["javascript"]=4
    ["c"]=5
    ["c++"]=6
    ["c#"]=7
    ["php"]=8
    ["kotlin"]=9
    ["swift"]=10
    ["java"]=11
    ["ruby"]=12
    ["haskell"]=13
    ["rust"]=14
    ["dart"]=15
    ["elixir"]=16
    ["elm"]=17
    ["julia"]=18
)
declare -A FRAMEWORK_ID_MAP=(
    ["fastapi"]=100
    ["django"]=101
    ["flask"]=102
    ["numpy"]=103
    ["pandas"]=104
    ["matplotlib"]=105
    ["pytorch"]=106
    ["seaborn"]=107
    ["tensorflow"]=108
    ["opencv"]=109
    ["sklearn"]=110
    ["nextjs"]=213
    ["nuxtjs"]=214
    ["react"]=215
    ["vue"]=216
    ["express"]=217
    ["nestjs"]=218
    ["jquery"]=219
    ["cakephp"]=300
    ["laravel"]=301
    ["symfony"]=302
    ["rubyonrails"]=400
    ["gin"]=500
    ["echo"]=501
)
declare -A TOOL_ID_MAP=(
    ["python-fastapi"]=100
    ["python-django"]=101
    ["python-flask"]=102
    ["python-numpy"]=103
    ["python-pandas"]=104
    ["python-matplotlib"]=105
    ["python-pytorch"]=106
    ["python-seaborn"]=107
    ["python-tensorflow"]=108
    ["python-opencv"]=109
    ["python-sklearn"]=110
    ["typescript-nextjs"]=213
    ["typescript-nuxtjs"]=214
    ["typescript-react"]=215
    ["typescript-vue"]=216
    ["typescript-express"]=217
    ["typescript-nestjs"]=218
    ["javascript-jquery"]=219
    ["php-cakephp"]=300
    ["php-laravel"]=301
    ["php-symfony"]=302
    ["ruby-rubyonrails"]=400
    ["go-gin"]=500
    ["go-echo"]=501
)
declare -A ALGORITHM_ID_MAP=(
    ["dfs"]=1
    ["bfs"]=2
    ["dp"]=3
    ["binary_search"]=4
    ["union_find"]=5
    ["bubble_sort"]=6
    ["quick_sort"]=7
    ["merge_sort"]=8
    ["selection_sort"]=9
    ["insertion_sort"]=10
    ["heap_sort"]=11
    ["stack"]=12
    ["queue"]=13
    ["gcd"]=14
    ["prime_factorization"]=15
)
declare -A PATTERN_ID_MAP=(
    ["abstract_factory"]=1
    ["builder"]=2
    ["factory_method"]=3
    ["prototype"]=4
    ["singleton"]=5
    ["adapter"]=6
    ["bridge"]=7
    ["composite"]=8
    ["decorator"]=9
    ["facade"]=10
    ["flyweight"]=11
    ["proxy"]=12
    ["chain_of_responsibility"]=13
    ["command"]=14
    ["interpreter"]=15
    ["iterator"]=16
    ["mediator"]=17
    ["momento"]=18
    ["observer"]=19
    ["state"]=20
    ["template_method"]=21
    ["strategy"]=22
    ["visitor"]=23
)

export LANGUAGES_ID_MAP
export FRAMEWORKS_ID_MAP
export TOOLS_ID_MAP
export ALGORITHMS_ID_MAP
export PATTERNS_ID_MAP
