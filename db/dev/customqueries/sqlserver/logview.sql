create view dbo.vw_log AS

SELECT [id]
      ,[created_at]
      ,[updated_at]
      ,[deleted_at]
      ,[user_id]
      ,[info]
      ,[func_name]
      ,[event_start_time]
      ,[event_end_time]
	  ,cast(DATEDIFF(millisecond, [event_start_time], [event_end_time]) as float) / 1000 AS diff_sec
      ,[critical_error]
      ,[error_message]
  FROM [dbo].[upm_loggers]
  --order by 1 desc