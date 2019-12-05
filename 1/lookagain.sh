# find -name '*.sh' | sed 's:.*/::g' | sed 's:.sh::g' 
# find -name '*.sh' 
# find -name '*.sh' | cut -f 2 
# find -name '*.sh' | cut -f 2 -d '.'
find -name '*.sh' | cut -d '.' -f 2 | sed "s:.*/::g"