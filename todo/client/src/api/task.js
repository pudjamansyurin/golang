import axios from "axios";

const api = axios.create({
  // timeout: 1000,
  baseURL: "http://localhost:8080/api",
  headers: { "Content-Type": "application/x-www-form-urlencoded" },
});

api.interceptors.response.use(
  (res) => {
    console.log(res);
    return res;
  },
  (err) => {
    console.error(err);
    throw err;
  }
);

const getTask = () => {
  return api.get(`/task`);
};

const createTask = (payload) => {
  return api.post(`/task`, { task: payload });
};

const updateTask = (id) => {
  return api.put(`/task/${id}`);
};

const undoTask = (id) => {
  return api.put(`/undoTask/${id}`);
};

const deleteTask = (id) => {
  return api.delete(`/deleteTask/${id}`);
};

export { getTask, createTask, updateTask, undoTask, deleteTask };
