
export function get(url, done, fail) {
    call("get", url, null, done, fail);
}

function call(method, url, data, done, fail) {

    if (!fail) {
        fail = function (e) {
            console.log(e);
        }
    }

    fetch(
        `${process.env.CONFIG.backend}/api${url}`,
        {
            method: method,
            data: data
        }
    )
        .then(res => res.json())
        .then(done).catch(fail);
}