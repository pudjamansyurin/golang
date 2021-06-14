const ls = window.localStorage;

const authProvider = {
  login: ({ username }) => {
    ls.setItem("username", username);
    return Promise.resolve();
  },
  logout: () => {
    ls.removeItem("username");
    return Promise.resolve();
  },
  checkError: ({ status }) => {
    if ([401, 403].includes(status)) {
      ls.removeItem("username");
      return Promise.reject();
    }
    return Promise.resolve();
  },
  checkAuth: () => {
    return ls.getItem("username") ? Promise.resolve() : Promise.reject();
  },
  getPermissions: () => Promise.resolve(),
};

export default authProvider;
