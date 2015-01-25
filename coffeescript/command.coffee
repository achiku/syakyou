{spawn} = require 'child_process'
ls = spawn 'ls', ['-l']

ls.stdout.on 'data', (data) -> console.log data.toString().trim()
