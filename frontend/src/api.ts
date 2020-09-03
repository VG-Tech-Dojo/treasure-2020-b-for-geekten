import {apiConfig} from "../env";

const toJson = async (res: Response) => {
  if (res.status == 404) {
    return "NotFound";
  }
  const js = await res.json();
  if (res.ok) {
    return js;
  } else {
    throw new Error(js.message);
  }
};

export const createOwner = async (idToken: string, body: any) => {
  const resp = await fetch(`${apiConfig.apiEndpoint}/owner`, {
    method: "POST",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`,
    }),
    credentials: "same-origin",
    body: JSON.stringify(body),
  });
  return await toJson(resp);
};

export const getOwnerMe = async (idToken: string) => {
  const resp = await fetch(`${apiConfig.apiEndpoint}/owner/me`, {
    method: "GET",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`,
    }),
    credentials: "same-origin",
  });
  return await toJson(resp);
};

export const getAllLesson = async (idToken: string) => {
  const resp = await fetch(`${apiConfig.apiEndpoint}/lesson`, {
    method: "GET",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`,
    }),
    credentials: "same-origin",
  });
  return await toJson(resp);
};

export const getLesson = async (lessonID: string) => {
  const resp = await fetch(`${apiConfig.apiEndpoint}/lesson/${lessonID}`, {
    method: "GET",
    credentials: "same-origin",
  });
  return await toJson(resp);
};

export const getResevedLesson = async (idToken: string) => {
  const resp = await fetch(`${apiConfig.apiEndpoint}/lesson`, {
    method: "GET",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`,
    }),
    credentials: "same-origin",
  });
  return await toJson(resp);
};

export const registLesson = async (idToken: string, body: any) => {
  const resp = await fetch(`${apiConfig.apiEndpoint}/lesson`, {
    method: "POST",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`,
    }),
    credentials: "same-origin",
    body: JSON.stringify(body),
  });

  if (resp.status === 500) {
    return null;
  }

  return await toJson(resp);
};

export const createWebsite = async (idToken: string, body: any) => {
  const resp = await fetch(`${apiConfig.apiEndpoint}/web`, {
    method: "POST",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`,
    }),
    credentials: "same-origin",
    body: JSON.stringify(body),
  });
  return await toJson(resp);
};

export const reserveLesson = async (lessonID: string, body: any) => {
  const resp = await fetch(`${apiConfig.apiEndpoint}/lesson/${lessonID}/reservation`, {
    method: "POST",
    credentials: "same-origin",
    body: JSON.stringify(body),
  });
  return await toJson(resp);
};

export const getOwnerLessons = async (owner_id: string) => {
  const resp = await fetch(`${apiConfig.apiEndpoint}/owner/${owner_id}/lesson`, {
    method: "GET",
    credentials: "same-origin",
  });
  return await toJson(resp);
};

export const getOwnerWebsite = async (owner_id: string) => {
  const resp = await fetch(`${apiConfig.apiEndpoint}/owner/${owner_id}/web`, {
    method: "GET",
    credentials: "same-origin",
  });
  return await toJson(resp);
};
