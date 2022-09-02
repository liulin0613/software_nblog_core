DROP TABLE IF EXISTS `docx`;

CREATE TABLE `docx` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `docx_id` varchar(27) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章id',
    `user_id` int(11) NOT NULL COMMENT '作者',
    `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
    `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
    `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章描述',
    `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示图片，默认为文本中的第一张',
    `tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标签',
    `dir_id` int(10) NOT NULL DEFAULT 0 COMMENT '所属文件夹 0.未归档',
    `is_top` int(2) NOT NULL DEFAULT 0 COMMENT '是否置顶',
    `is_delete` int(2) NOT NULL DEFAULT 0 COMMENT '是否删除',
    `scope` int(2) NOT NULL COMMENT '文章状态 1.公开 2.私密 3.粉丝可见',
    `create_time` datetime NOT NULL COMMENT '创建时间',
    `update_time` datetime NOT NULL COMMENT '修改时间',
    `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示图片，默认为文本中的第一张',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

DROP TABLE IF EXISTS `dir`;

CREATE TABLE `dir` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `dir_id` varchar(27) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件夹id',
    `user_id` int(11) NOT NULL COMMENT '作者',
    `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件夹名字',
    `is_delete` int(2) NOT NULL DEFAULT 0 COMMENT '是否删除',
    `create_time` datetime NOT NULL COMMENT '创建时间',
    `update_time` datetime NOT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10000 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;