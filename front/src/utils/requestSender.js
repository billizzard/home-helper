async function requestSender (url, data, options) {
    const backendUrl = window.appConfig.backendUrl;

    const settings = {
        method: options?.method || 'GET',
        setAlert: options?.setAlert || null,
        isFile: options?.isFile || false
    }

    const response = {
        data: null,
        error: null,
        redirect: null,
        statusCode: 200
    }

    const baseUrl = `${backendUrl}/api/v1`

    try {
        const res = await fetch(baseUrl + url, {
            method: settings.method,
            body: data
        })

        response.statusCode = res.status

        if (res.status == 200) {
            if (settings.isFile) {
                response.data = await res.blob();
                let url = window.URL.createObjectURL(response.data);
                let a = document.createElement('a');
                a.href = url;
                a.download = '';
                a.click();
                return response
            } else {
                response.data = await res.json();
                return response;
            }
        }

        response.statusCode = res.status

        process404(res.status, res)
        await process400(res.status, res, settings.setAlert)
    } catch (error) {
        response.error = {
            type: 'warning',
            message: 'Server is not responding'
        }

        response.data = null
        if (settings.setAlert != null) {
            settings.setAlert({type: response.error.type, message: response.error.message})
        }
    }

    return response
}

const process404 = (status, response) => {
    if (status == 404) {
        response.redirect = '/404'
    }

    return response
}

const process400 = async (status, res, setAlert) => {
    if (status == 400 && setAlert != null) {
        setAlert({type: 'warning', message: await res.text()})
    }
}

export default requestSender