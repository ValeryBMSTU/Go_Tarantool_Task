#!/usr/bin/env tarantool
-- настроить базу данных
box.cfg {
    listen = 3301
 }
 box.once('tester', function()
    s = box.schema.space.create('test')
    s:format({
        {name = 'key', type = 'string'},
        {name = 'value', type = 'string'}
        })
    s:create_index('primary', {
        type = 'hash',
        parts = {'key'}
        })

    box.schema.user.grant('guest', 'read,write,execute', 'universe')


    print('!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!')
end)
