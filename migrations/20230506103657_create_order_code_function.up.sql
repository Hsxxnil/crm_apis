create function order_code()
    returns text as
    $$
declare
old_id varchar :=(select  code  from orders order by code desc limit 1);
    id_number char(3) :='001';
    datetime char(4) :=substring(cast(now() as varchar),3,2)||substring(cast(now() as varchar),6,2);
    new_id varchar ;
    num integer;
begin
    if old_id is null then
        new_id:='O'||datetime||id_number;
return new_id;
end if;

    if datetime=substring(old_id,2,4) then
        num :=cast(right(old_id,3) as integer)+1;
        id_number:=
        case
            when num<10 then '00'||num
            when num<100 then '0'||num
            when num<1000 then cast(num as varchar)
end;
end if;

    new_id:='O'||datetime||id_number;
return new_id;
end;
$$
language 'plpgsql';