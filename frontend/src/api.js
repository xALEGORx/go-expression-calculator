import axios from 'axios';

const API_URL = 'http://127.0.0.1:8000/api/v1/';

class ApiService {
  getTasks() {
    return axios.get(API_URL + "task");
  }
}

export default new ApiService();
