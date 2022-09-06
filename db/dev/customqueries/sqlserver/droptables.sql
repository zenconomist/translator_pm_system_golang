declare @name nvarchar(100), @sql nvarchar(max)

declare c cursor for
select [name] from sys.tables
where [name] not in ('upm_configs','upm_loggers', 'distinct_types', 'dto_configs')

open c
fetch next from c into @name

while @@FETCH_STATUS = 0
begin

	set @sql = 'drop table if exists ' + @name 
	exec (@sql)
	fetch next from c into @name

end

close c
deallocate c

