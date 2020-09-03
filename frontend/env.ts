export const zoomConfig = {
    zoomRedirectURI: process.env.ZOOM_REDIRECT_URI || "",
    zoomClientID: process.env.ZOOM_CLIENT_ID || ""
};

export const apiConfig = {
    apiEndpoint: process.env.BACKEND_API_BASE || ""
};

export const payjpConfig = {
    payjpPkTest: process.env.PAYJP_PK_TEST || ""
};

export const initEnv = () => {
    if (zoomConfig.zoomRedirectURI === "") {
        console.log("failed to get env ZOOM_REDIRECT_URI");
    }
    if (zoomConfig.zoomClientID === "") {
        console.log("failed to get env ZOOM_CLIENT_ID");
    }
    if (apiConfig.apiEndpoint === "") {
        console.log("failed to get env BACKEND_API_BASE");
    }
    if (payjpConfig.payjpPkTest === "") {
        console.log("failed to get env PAYJP_PK_TEST");
    }
};

