function! Hello()
    echo "Hello, world!"
    let flag = "hogehoge"
    echo exists('flag')
    if exists('flag')
        echo 'flag exists'
    endif
    echo has('flag')
    if !has('flag')
        echo 'flag exists'
    endif
endfunction

call Hello()
