async function request(input: string, init?: RequestInit): Promise<Response> {
    const [method, url] = input.split(' ');
    const token = localStorage.getItem('token');
    let headers: string[][] = [];
    if (token) {
        headers.push(['Authorization', `Bearer ${token}`]);
    }
    return fetch(url, {
        method,
        headers,
        ...init
    });
}

export default {
    request
};
