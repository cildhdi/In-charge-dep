import utils from './utils';
import { UserInfo } from './types';

export const Urls = {
    login: "POST /api/login",
    register: "POST /api/register",
    sendCode: "POST /api/send-code",
    adminRegister: "POST /api/admin-register",
    allUsers: "GET /api/user/all",
    reachable: "GET /api/auth/reachable",
    unreachable: "GET /api/auth/unreachable"
};


export const Apis = {
    sendCode: async (phone: string) => {
        let response = await utils.request(Urls.sendCode, {
            body: JSON.stringify({
                phone
            })
        });
        if (response.ok && (await response.json()).code != 0) {
            throw Error("验证码发送失败");
        }
    },
    login: async (phone: string, code: string): Promise<string> => {
        let response = await utils.request(Urls.login, {
            body: JSON.stringify({
                phone,
                code
            })
        });
        if (response.ok) {
            let body = await response.json();
            if (body.code == 200) {
                localStorage.setItem('token', body.token);
                return body.token as string;
            } else {
                throw Error("登录失败");
            }
        } else {
            throw Error("登录失败");
        }
    },
    allUsers: async (): Promise<UserInfo[]> => {
        let response = await utils.request(Urls.allUsers);
        if (response.ok) {
            let body = await response.json();
            console.log(body);
            if (body.code == 0 && body.data) {
                return body.data as UserInfo[];
            } else {
                throw Error("获取用户失败");
            }
        } else {
            throw Error("获取用户失败");
        }
    },
    adminRegister: async (phone: string, role: 0 | 1 | 2, name: string) => {
        let response = await utils.request(Urls.adminRegister, {
            body: JSON.stringify({
                phone,
                role,
                name
            })
        });
        if (response.ok) {
            let body = await response.json();
            if (body.code != 0) {
                throw Error(body.msg);
            }
        }
    }
}
