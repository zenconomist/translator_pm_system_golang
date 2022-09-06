;with config as (
	select config_id, folder_type, hierarchy, folder_name , parent_folder_config_id 
	from share_point_folder_configs spfc 
), pandt as (
	select project_id, id, concat(id, '-', case when task_type = '' then 'translate' else task_type end) as taskn , case when task_type = '' then 'translate' else task_type end as task_type from tasks where deleted_at is null
), pr as (
	select id from projects p where deleted_at is null
), base0 as (
	select id as project_id, 0 as task_id, c.folder_type, c.hierarchy
	,case when c.folder_name='<<project>>' then id::text
		else folder_name 
	end as folder_name
	, c.config_id, c.parent_folder_config_id
	from pr
	join config c
		on c.folder_type = 'pf'	
	union all
	select t.project_id as project_id, t.id as task_id, c.folder_type, c.hierarchy
	,case when c.folder_name='<<project>>' then id::text
		when c.folder_name = '<<task>>' then t.taskn
		when c.folder_name = '<<taskType>>' then t.task_type
		else folder_name 
	end as folder_name
	, c.config_id, c.parent_folder_config_id
	from pandt t
	left join config c
		on c.folder_type <> 'pf'
	except 
	select project_id, task_id, folder_type, hierarchy, folder_name , config_id, parent_folder_config_id
	from share_point_folders spf 
	)
select *
from base0
order by project_id, task_id, config_id asc
;