import axios from "axios";

export const useAxiosCreate = () => {
  const instance = axios.create({
    baseURL: "http://localhost:8080",
  });
  return instance;
};
