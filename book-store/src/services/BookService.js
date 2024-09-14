import axios from 'axios';

const API_URL = 'http://localhost:9010/book/';

export default {
  getAllBooks() {
    return axios.get(API_URL + 'all');
  },

  getBookById(bookId) {
    return axios.get(`${API_URL}${bookId}`);
  },

  createBook(book) {
    return axios.post(API_URL + 'create', book);
  },

  updateBook(bookId, book) {
    return axios.put(`${API_URL}update/${bookId}`, book);
  },

  deleteBook(bookId) {
    return axios.delete(`${API_URL}delete/${bookId}`);
  }
};
