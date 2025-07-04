CREATE TABLE accounts (
        "id" bigserial PRIMARY KEY,
        "owner" varchar NOT NULL,
        "balance" bigint NOT NULL,
        "currency" varchar NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );


CREATE TABLE
    "transfers" (
        "id" bigserial PRIMARY KEY,
        "from_account_id" bigint NOT NULL,
        "to_account_id" bigint NOT NULL,
        "amount" bigint NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

ALTER TABLE "transfers"
ADD
    FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
ADD
    FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");