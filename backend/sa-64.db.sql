BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "users" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "orders" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"total_price"	real,
	"owner_id"	integer,
	CONSTRAINT "fk_users_order" FOREIGN KEY("owner_id") REFERENCES "users"("id"),
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "deliver_types" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"type"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "payment_methods" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"method"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "payments" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"pice"	real,
	"payment_time"	datetime,
	"order_id"	integer,
	"payment_method_id"	integer,
	"delivery_type_id"	integer,
	CONSTRAINT "fk_payment_methods_payment" FOREIGN KEY("payment_method_id") REFERENCES "payment_methods"("id"),
	CONSTRAINT "fk_deliver_types_payment" FOREIGN KEY("delivery_type_id") REFERENCES "deliver_types"("id"),
	CONSTRAINT "fk_orders_payment" FOREIGN KEY("order_id") REFERENCES "orders"("id"),
	PRIMARY KEY("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_users_email" ON "users" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_orders_deleted_at" ON "orders" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_deliver_types_deleted_at" ON "deliver_types" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_payment_methods_deleted_at" ON "payment_methods" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_payments_deleted_at" ON "payments" (
	"deleted_at"
);
COMMIT;
