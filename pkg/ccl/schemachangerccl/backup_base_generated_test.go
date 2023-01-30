// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

// Code generated by sctestgen, DO NOT EDIT.

package schemachangerccl

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/sctest"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
)

func TestBackup_base_add_column(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/add_column", newCluster)
}
func TestBackup_base_add_column_default_seq(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/add_column_default_seq", newCluster)
}
func TestBackup_base_add_column_default_unique(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/add_column_default_unique", newCluster)
}
func TestBackup_base_add_column_no_default(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/add_column_no_default", newCluster)
}
func TestBackup_base_alter_table_add_check_vanilla(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_check_vanilla", newCluster)
}
func TestBackup_base_alter_table_add_check_with_seq_and_udt(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_check_with_seq_and_udt", newCluster)
}
func TestBackup_base_alter_table_add_foreign_key(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_foreign_key", newCluster)
}
func TestBackup_base_alter_table_add_primary_key_drop_rowid(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_primary_key_drop_rowid", newCluster)
}
func TestBackup_base_alter_table_add_unique_without_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_unique_without_index", newCluster)
}
func TestBackup_base_alter_table_alter_column_set_not_null(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_alter_column_set_not_null", newCluster)
}
func TestBackup_base_alter_table_alter_primary_key_drop_rowid(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_alter_primary_key_drop_rowid", newCluster)
}
func TestBackup_base_alter_table_alter_primary_key_vanilla(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_alter_primary_key_vanilla", newCluster)
}
func TestBackup_base_alter_table_drop_constraint_check(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_drop_constraint_check", newCluster)
}
func TestBackup_base_alter_table_drop_constraint_fk(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_drop_constraint_fk", newCluster)
}
func TestBackup_base_alter_table_drop_constraint_uwi(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_drop_constraint_uwi", newCluster)
}
func TestBackup_base_create_function(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/create_function", newCluster)
}
func TestBackup_base_create_function_in_txn(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/create_function_in_txn", newCluster)
}
func TestBackup_base_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/create_index", newCluster)
}
func TestBackup_base_drop_column_basic(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_basic", newCluster)
}
func TestBackup_base_drop_column_computed_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_computed_index", newCluster)
}
func TestBackup_base_drop_column_create_index_separate_statements(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_create_index_separate_statements", newCluster)
}
func TestBackup_base_drop_column_unique_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_unique_index", newCluster)
}
func TestBackup_base_drop_column_with_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_with_index", newCluster)
}
func TestBackup_base_drop_function(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_function", newCluster)
}
func TestBackup_base_drop_index_hash_sharded_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_index_hash_sharded_index", newCluster)
}
func TestBackup_base_drop_index_partial_expression_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_index_partial_expression_index", newCluster)
}
func TestBackup_base_drop_index_vanilla_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_index_vanilla_index", newCluster)
}
func TestBackup_base_drop_index_with_materialized_view_dep(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_index_with_materialized_view_dep", newCluster)
}
func TestBackup_base_drop_multiple_columns_separate_statements(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_multiple_columns_separate_statements", newCluster)
}
func TestBackup_base_drop_schema(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_schema", newCluster)
}
func TestBackup_base_drop_table(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.Backup(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_table", newCluster)
}
