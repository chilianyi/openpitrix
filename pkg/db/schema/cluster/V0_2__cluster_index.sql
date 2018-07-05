CREATE INDEX cluster_name_index ON cluster (name ASC);
CREATE INDEX cluster_app_id_index ON cluster (app_id ASC);
CREATE INDEX cluster_version_id_index ON cluster (version_id ASC);
CREATE INDEX cluster_runtime_id_index ON cluster (runtime_id ASC);
CREATE INDEX cluster_node_name_index ON cluster_node (name ASC);
CREATE INDEX cluster_node_instance_id_index ON cluster_node (instance_id ASC);
CREATE INDEX cluster_node_volume_id_index ON cluster_node (volume_id ASC);
CREATE INDEX cluster_node_private_ip_index ON cluster_node (private_ip ASC);
CREATE INDEX cluster_node_role_index ON cluster_node (role ASC);