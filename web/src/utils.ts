async function request(input: string, init?: RequestInit): Promise<Response> {
    const [method, url] = input.split(' ');
    return fetch(url, {
        method,
        ...init
    });
}

export default {
    request
};