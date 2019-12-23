import React, { useEffect, useState } from 'react';
import { Table, message, Button, Modal } from 'antd';
import Moment from 'moment';

import { Apis } from '../../api';
import { UserInfo } from '../../types';

const columns = [
    {
        title: "注册时间",
        dataIndex: "CreatedAt",
        key: "created_at",
        render: (time: string) => {
            return Moment(time).toLocaleString();
        }
    },
    {
        title: "名称",
        dataIndex: "Name",
        key: "name"
    },
    {
        title: "电话号码",
        dataIndex: "Phone",
        key: "phone"
    },
    {
        title: "是否封禁",
        dataIndex: "Banned",
        key: "banned",
        render: (banned: number) => {
            return banned == 0 ? "否" : "是";
        }
    },
    {
        title: "操作",
        key: "action",
        render: (text: string, user: UserInfo) => {
            return <div>
                <Button type='dashed'>修改电话号码</Button>
                <div style={{ margin: 5, display: 'inline' }} />
                <Button type='danger'>封禁</Button>
            </div>
        }
    }
];
const MerchantManage: React.FC = (props) => {

    let [users, setUsers] = useState([] as UserInfo[]);
    const fetchUsersKey = "fetchUsersKey";
    const fetchUsers = async () => {
        try {
            message.loading({
                content: "获取用户中",
                key: fetchUsersKey,
            });
            let users = await Apis.allUsers();
            setUsers(users);
            message.success({
                content: "获取用户成功",
                key: fetchUsersKey,
            });
        } catch (e) {
            message.error({
                content: `获取用户失败 ${e}`,
                key: fetchUsersKey,
            });
        }
    };

    const registerUser = () => {
        Modal.confirm({
            title: "注册",
            content: "form items",
            okText: "确认",
            cancelText: "取消",
        });
    };

    useEffect(() => {
        fetchUsers();
    }, []);

    return (
        <div style={{
            display: 'flex',
            alignItems: 'center',
            flexDirection: 'column',
            width: '100%'
        }}>
            <div style={{ width: '100%' }}>
                <Button type='default' onClick={fetchUsers}>刷新</Button>
                <div style={{ margin: 5, display: 'inline' }} />
                <Button type='default' onClick={registerUser}>商户注册</Button>
            </div>
            <Table columns={columns} dataSource={users} style={{
                width: '100%'
            }} />

        </div>
    );
}

export default MerchantManage;
