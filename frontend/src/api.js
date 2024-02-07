import axios from 'axios';

const API_URL = 'http://127.0.0.1:8000/api/v1/';

class ApiService {
  getTasks() {
    return axios.get(API_URL + "task");
  }
  addTask(expression) {
    return axios.post(API_URL + "task", {expression: expression})
  }
}

const apiService = new ApiService();

export default apiService;
