-- INSERT
insert into task_histories (
    original_id
    ,created_at
    -- ,updated_at
    -- ,deleted_at
    ,flg_is_current
    ,flg_is_deleted
    ,dat_from
    ,dat_to
    ,project_id
    ,batch_id
    ,order_within_batch
    ,task_type
    ,project_manager
    ,task_state_id
    ,task_time_state_id
    ,source_lang
    ,target_lang
    ,task_speciality
    ,prep_disabled
    ,review_disabled
    ,prep_billable
    ,review_billable
    ,preparer_id
    ,reviewer_id
    ,customer_due_date
    ,cust_quantity
    ,customer_price
    ,customer_price_type_id
    ,customer_unit_type_id
    ,customer_unique_price
    ,supplier_id
    ,supplier_due_date
    ,supplier_price_type_id
    ,supplier_unit_type_id
    ,supplier_quantity
    ,supplier_to_bill
    ,supplier_time_state_id
    ,billed_by_supplier
    ,billing_fulfillment_date
    ,billing_invoice_number
    ,supplier_unique_price
    ,prepare_state_id
    ,review_state_id
)
select
id
,now() as created_at
,true as flg_is_current
,false as flg_is_deleted
,now() as dat_from
,'2099-12-31T23:59:59+00:00' as dat_to
,project_id
,batch_id
,order_within_batch
,task_type
,project_manager
,task_state_id
,task_time_state_id
,source_lang
,target_lang
,task_speciality
,prep_disabled
,review_disabled
,prep_billable
,review_billable
,preparer_id
,reviewer_id
,tcp_customer_due_date
,tcp_cust_quantity
,tcp_customer_price
,tcp_customer_price_type_id
,tcp_customer_unit_type_id
,tcp_customer_unique_price
,tsp_supplier_id
,tsp_supplier_due_date
,tsp_supplier_price_type_id
,tsp_supplier_unit_type_id
,tsp_supplier_quantity
,tsp_supplier_to_bill
,tsp_supplier_time_state_id
,tsp_billed_by_supplier
,tsp_billing_fulfillment_date
,tsp_billing_invoice_number
,tsp_supplier_unique_price
,prepare_state_id
,review_state_id
from tasks
where id in (
    select id from tasks 
    where deleted_at is null
    except
    select original_id from task_histories where flg_is_current = true and flg_is_deleted = false
);


-- UPDATE

DROP TABLE IF EXISTS to_be_updated;

CREATE TEMPORARY TABLE to_be_updated AS
SELECT id
,project_id
,batch_id
,order_within_batch
,task_type
,project_manager
,task_state_id
,task_time_state_id
,source_lang
,target_lang
,task_speciality
,prep_disabled
,review_disabled
,prep_billable
,review_billable
,preparer_id
,reviewer_id
,tcp_customer_due_date
,tcp_cust_quantity
,tcp_customer_price
,tcp_customer_price_type_id
,tcp_customer_unit_type_id
,tcp_customer_unique_price
,tsp_supplier_id
,tsp_supplier_due_date
,tsp_supplier_price_type_id
,tsp_supplier_unit_type_id
,tsp_supplier_quantity
,tsp_supplier_to_bill
,tsp_supplier_time_state_id
,tsp_billed_by_supplier
,tsp_billing_fulfillment_date
,tsp_billing_invoice_number
,tsp_supplier_unique_price
,prepare_state_id
,review_state_id
FROM tasks WHERE 1=2;

INSERT INTO to_be_updated
select
id
,project_id
,batch_id
,order_within_batch
,task_type
,project_manager
,task_state_id
,task_time_state_id
,source_lang
,target_lang
,task_speciality
,prep_disabled
,review_disabled
,prep_billable
,review_billable
,preparer_id
,reviewer_id
,tcp_customer_due_date
,tcp_cust_quantity
,tcp_customer_price
,tcp_customer_price_type_id
,tcp_customer_unit_type_id
,tcp_customer_unique_price
,tsp_supplier_id
,tsp_supplier_due_date
,tsp_supplier_price_type_id
,tsp_supplier_unit_type_id
,tsp_supplier_quantity
,tsp_supplier_to_bill
,tsp_supplier_time_state_id
,tsp_billed_by_supplier
,tsp_billing_fulfillment_date
,tsp_billing_invoice_number
,tsp_supplier_unique_price
,prepare_state_id
,review_state_id
from tasks
where id in (
    select id from tasks 
    where deleted_at is null
    intersect
    select original_id from task_histories where flg_is_current = true and flg_is_deleted = false
)

except

select
original_id
,project_id
,batch_id
,order_within_batch
,task_type
,project_manager
,task_state_id
,task_time_state_id
,source_lang
,target_lang
,task_speciality
,prep_disabled
,review_disabled
,prep_billable
,review_billable
,preparer_id
,reviewer_id
,customer_due_date
,cust_quantity
,customer_price
,customer_price_type_id
,customer_unit_type_id
,customer_unique_price
,supplier_id
,supplier_due_date
,supplier_price_type_id
,supplier_unit_type_id
,supplier_quantity
,supplier_to_bill
,supplier_time_state_id
,billed_by_supplier
,billing_fulfillment_date
,billing_invoice_number
,supplier_unique_price
,prepare_state_id
,review_state_id
from task_histories
where flg_is_current = true and flg_is_deleted = false
;


update task_histories
set flg_is_current = false, dat_to = now()
from task_histories th
inner join to_be_updated u
    on th.original_id=u.id
    where th.flg_is_current = true and th.flg_is_deleted = false;

insert into task_histories (
    original_id
    ,created_at
    -- ,updated_at
    -- ,deleted_at
    ,flg_is_current
    ,flg_is_deleted
    ,dat_from
    ,dat_to
    ,project_id
    ,batch_id
    ,order_within_batch
    ,task_type
    ,project_manager
    ,task_state_id
    ,task_time_state_id
    ,source_lang
    ,target_lang
    ,task_speciality
    ,prep_disabled
    ,review_disabled
    ,prep_billable
    ,review_billable
    ,preparer_id
    ,reviewer_id
    ,customer_due_date
    ,cust_quantity
    ,customer_price
    ,customer_price_type_id
    ,customer_unit_type_id
    ,customer_unique_price
    ,supplier_id
    ,supplier_due_date
    ,supplier_price_type_id
    ,supplier_unit_type_id
    ,supplier_quantity
    ,supplier_to_bill
    ,supplier_time_state_id
    ,billed_by_supplier
    ,billing_fulfillment_date
    ,billing_invoice_number
    ,supplier_unique_price
    ,prepare_state_id
    ,review_state_id
)
select
id
,now() as created_at
,true as flg_is_current
,false as flg_is_deleted
,now() as dat_from
,'2099-12-31T23:59:59+00:00' as dat_to
,project_id
,batch_id
,order_within_batch
,task_type
,project_manager
,task_state_id
,task_time_state_id
,source_lang
,target_lang
,task_speciality
,prep_disabled
,review_disabled
,prep_billable
,review_billable
,preparer_id
,reviewer_id
,tcp_customer_due_date
,tcp_cust_quantity
,tcp_customer_price
,tcp_customer_price_type_id
,tcp_customer_unit_type_id
,tcp_customer_unique_price
,tsp_supplier_id
,tsp_supplier_due_date
,tsp_supplier_price_type_id
,tsp_supplier_unit_type_id
,tsp_supplier_quantity
,tsp_supplier_to_bill
,tsp_supplier_time_state_id
,tsp_billed_by_supplier
,tsp_billing_fulfillment_date
,tsp_billing_invoice_number
,tsp_supplier_unique_price
,prepare_state_id
,review_state_id
from to_be_updated;


update task_histories
set flg_is_current = false, flg_is_deleted = true, dat_to = now()
where original_id in (
    select original_id from task_histories where flg_is_current = true and flg_is_deleted = false
    except
    select id from tasks 
    where deleted_at is null
);