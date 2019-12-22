import React, { useState } from 'react';
import { Button, Input, message } from 'antd';

import { Apis } from '../api';


export default function () {
  const [phone, setPhone] = useState("");
  const [code, setCode] = useState("");

  const itemStyle: React.CSSProperties = {
    width: 400,
    marginTop: 5,
    marginBottom: 5
  };

  const onSendCode = async () => {
    if (phone.length != 11) {
      message.warn("手机号格式错误");
      return;
    }
    const sendCodeStatusKey = 'sendCodeStatusKey';
    message.loading({
      content: "验证码发送中...",
      key: sendCodeStatusKey
    });
    try {
      await Apis.sendCode(phone);
      message.success({
        content: "验证码发送成功",
        key: sendCodeStatusKey
      });
    } catch (e) {
      console.log(e);
      message.error({
        content: "验证码发送失败",
        key: sendCodeStatusKey
      });
    }
  }

  const onLogin = async () => {
    if (phone.length != 11 || code.length != 4) {
      message.warn("手机号或验证码格式错误");
      return;
    }

    const loginStatusKey = 'loginStatusKey';

    message.loading({
      content: "登录中...",
      key: loginStatusKey
    });
    try {
      await Apis.login(phone, code);
      message.success({
        content: "登录成功",
        key: loginStatusKey
      });
    } catch (e) {
      console.log(e);
      message.error({
        content: "登录失败，身份验证错误",
        key: loginStatusKey
      });
    }
  }

  return (
    <div style={{
      display: 'flex',
      flexDirection: 'column',
      justifyContent: 'center',
      alignItems: 'center',
      height: '100%',
      width: '100%',
      backgroundColor: '#eee'
    }}>
      <h1>智联平台</h1>
      <h2>后台登录</h2>
      <div style={{
        ...itemStyle,
        display: 'flex',
        flexDirection: 'row'
      }}>
        <Input value={phone} placeholder={"请输入手机号"} onChange={
          (e) => setPhone(e.target.value)
        } />
        <Button style={{
          marginLeft: 5,
        }} onClick={onSendCode}>获取验证码</Button>
      </div>
      <Input placeholder={"请输入验证码"} style={itemStyle} value={code} onChange={
        (e) => setCode(e.target.value)
      } />
      <Button type='primary' block={true} style={itemStyle} onClick={onLogin}>
        登录
      </Button>
    </div>
  );
}
