create function split_array(arr text[])
    returns text[] as
    $$
declare
res TEXT[] := '{}';
begin
    for i in 1..array_length(arr, 1) loop
        for j in 1..length(arr[i]) loop
            res := array_append(res, substring(arr[i] from j for 1));
        end loop;
    end loop;

return res;
end;
$$ language 'plpgsql';

