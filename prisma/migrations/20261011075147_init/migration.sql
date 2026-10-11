-- CreateEnum
CREATE TYPE "property_member_role" AS ENUM ('owner', 'manager', 'tenant');

-- CreateTable
CREATE TABLE "users" (
    "id" SERIAL NOT NULL,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "properties" (
    "id" SERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,

    CONSTRAINT "properties_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "property_members" (
    "property_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "role" "property_member_role" NOT NULL DEFAULT 'tenant',

    CONSTRAINT "property_members_pkey" PRIMARY KEY ("property_id","user_id")
);

-- AddForeignKey
ALTER TABLE "property_members" ADD CONSTRAINT "property_members_property_id_fkey" FOREIGN KEY ("property_id") REFERENCES "properties"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "property_members" ADD CONSTRAINT "property_members_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
