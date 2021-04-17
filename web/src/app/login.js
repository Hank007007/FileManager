import ReactDOM from "react-dom";
import "./styles.scss";
import "./manifest.json"
import "./antd.min.css"
import App from "./app";
import LikeButton from "./components/like-button";
import Header from "./components/header";
import cfg from '../config';
import { storage } from 'ra-lib';
import React from "react";


var mountNode = document.getElementById("app");
ReactDOM.render(<App />, mountNode);

document.querySelectorAll(".like-button-component").forEach(domContainer => {
    ReactDOM.render(<LikeButton {...domContainer.dataset} />, domContainer);
});
document.querySelectorAll(".header-component").forEach(domContainer => {
    ReactDOM.render(<Header {...domContainer.dataset} />, domContainer);
});

const LOGIN_USER_STORAGE_KEY = 'login-user';
const {baseName} = cfg;
const sessionStorage = window.sessionStorage;

/**
 * nav browser，with baseName
 * @param href
 * @returns {string|*}
 */
export function locationHref(href) {
    if (href.startsWith('http')) return window.location.href = href;

    return window.location.href = `${baseName}${href}`;
}

/**
 * verify user permission
 * @param code
 */
export function hasPermission(code) {
    const loginUser = getLoginUser();
    return loginUser.permissions.includes(code);
}

/**
 * get login user information
 * @returns {any}
 */
export function getLoginUser() {
    const loginUser = sessionStorage.getItem(LOGIN_USER_STORAGE_KEY);

    return loginUser ? JSON.parse(loginUser) : null;
}

/**
 * set user info
 * @param loginUser current login user
 */
export function setLoginUser(loginUser = {}) {
    const {id, name, avatar, token, permissions, ...others} = loginUser;
    const userStr = JSON.stringify({
        id,             // 用户id 必须
        name,           // 用户名 必须
        avatar,         // 用头像 非必须
        token,          // 登录凭证 非必须 ajax请求有可能会用到，也许是cookie
        permissions,    // 用户权限
        ...others,      // 其他属性
    });

    sessionStorage.setItem(LOGIN_USER_STORAGE_KEY, userStr);
}

/**
 * validate whether user is login based on user info existed
 * @returns {boolean}
 */
export function isLogin() {
    return !!getLoginUser();
}

/**
 * go to home page
 */
export function toHome() {
    // 跳转页面，优先跳转上次登出页面
    const lastHref = window.sessionStorage.getItem('last-href');

    locationHref(lastHref || '/');
}

/**
 * nav to login page
 */
export function toLogin() {
    const loginPath = '/login';

    // verify whether current page is login page, if yes, return, to avoid the death loop
    const pathname = window.location.pathname;
    const isLogin = pathname.indexOf(loginPath) !== -1;

    if (isLogin) return null;

    // clear the data
    storage.session.clear();
    sessionStorage.clear();
    sessionStorage.setItem('last-href', window.location.pathname);

    locationHref(loginPath);

    return null;
}
