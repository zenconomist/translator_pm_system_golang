CREATE OR REPLACE PROCEDURE public.usp_task_state_change_load()
 LANGUAGE plpgsql
AS $procedure$
DECLARE _config_id int;
DECLARE _now timestamptz = NOW();
BEGIN

    --add config head with validity info
    INSERT INTO task_state_config_heads (created_at, active_from, active_to) values (_now, _now, '2999-12-31T00:00:000000')
    RETURNING ID INTO _config_id

    --invalidating all older state config id's
    ;UPDATE task_state_config_heads
    SET active_to = _now, updated_at = _now
    WHERE active_to = '2999-12-31T00:00:00+00:00'
    AND id <> _config_id
    
    --adding states to state table
    ;INSERT INTO task_states (created_at, task_state_config_head_id, task_state_code, task_state_name, task_state_order)
    VALUES 
    (_now, _config_id, 'op', 'Open',         1),
    (_now, _config_id, 'of', 'Offered',      2),
    (_now, _config_id, 'ip', 'InProgress',   3),
    (_now, _config_id, 're', 'Ready',        4),
    (_now, _config_id, 'dl', 'Delivered',    5),
    (_now, _config_id, 'bl', 'Billed',       6),
    (_now, _config_id, 'pd', 'Pending',      7),
    (_now, _config_id, 'cl', 'Claimed',      8),
    (_now, _config_id, 'ah', 'Archived',     9)

    --inserting descartes multiplication of states into state changes
    ;INSERT INTO task_state_changes (created_at, from_task_state_code, to_task_state_code, from_task_state_id, to_task_state_id, is_allowed, state_change_info)
    SELECT _now, ts.from_tsc, ts.to_tsc, ts.from_task_state_id, ts.to_task_state_id, ts.is_allowed, ts.sci
    FROM (
            SELECT 
                ts1.task_state_code AS from_tsc, ts2.task_state_code AS to_tsc
                ,ts1.id as from_task_state_id, ts2.id as to_task_state_id
                ,false as is_allowed
                ,'unknown' as sci
                , ts1.task_state_config_head_id
            FROM task_states ts1
            CROSS JOIN task_states ts2
            WHERE ts1.task_state_config_head_id = _config_id
            AND ts2.task_state_config_head_id = _config_id
    ) ts
    INNER JOIN task_state_config_heads tsch
        ON ts.task_state_config_head_id=tsch.id
        AND ts.task_state_config_head_id=_config_id

    --update if id's are the same -> state can not change
    ;UPDATE task_state_changes
    SET state_change_info = 'state can not change to itself'
    where from_task_state_id = to_task_state_id

  
    
    ;update task_state_changes tsc
    SET is_allowed = _asc.is_allowed, state_change_info = _asc.state_change_info
    FROM  task_states ts1, task_states ts2, actual_state_changes _asc
    WHERE ts1.id=tsc.from_task_state_id
    AND ts2.id=tsc.to_task_state_id
    AND tsc.from_task_state_code =_asc."from"
    AND tsc.to_task_state_code =_asc."to"
    AND ts1.task_state_config_head_id = _config_id
    AND ts2.task_state_config_head_id = _config_id
        ;
/*
    --following steps:
        -- need to adjust state_changes table with message and error fields
        -- need to adjust task_states with primary key on task_state_config_head_id and task_state_code -> ensuring no duplicates
*/
    --call usp_task_state_change_testing();

end; $procedure$
;
