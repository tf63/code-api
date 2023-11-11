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
    ["none"]=1
)
declare -A TOOL_ID_MAP=(
    ["python-none"]=1
    ["go-none"]=2
    ["typescript-none"]=3
    ["javascript-none"]=4
    ["c-none"]=5
    ["c++-none"]=6
    ["c#-none"]=7
    ["php-none"]=8
    ["kotlin-none"]=9
    ["swift-none"]=10
    ["java-none"]=11
    ["ruby-none"]=12
    ["haskell-none"]=13
    ["rust-none"]=14
    ["dart-none"]=15
    ["elixir-none"]=16
    ["elm-none"]=17
    ["julia-none"]=18
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
