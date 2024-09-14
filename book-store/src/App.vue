<template>
  <div id="app">
    <h1>Book Store</h1>

    <!-- Create Book Form -->
    <div>
      <h2>Create a New Book</h2>
      <form @submit.prevent="createBook">
        <input v-model="newBook.name" placeholder="Book Name" required />
        <input v-model="newBook.author" placeholder="Author" required />
        <input v-model="newBook.publication" placeholder="Publication" required />
        <button type="submit" :disabled="isCreating">Add Book</button>
      </form>
    </div>

    <!-- Book List -->
    <div>
      <h2>Books</h2>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Author</th>
            <th>Publication</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="book in books" :key="book.ID">
            <td>{{ book.ID }}</td>
            <td>{{ book.name }}</td>
            <td>{{ book.author }}</td>
            <td>{{ book.publication }}</td>
            <td>
              <button @click="confirmDelete(book.ID)">Delete</button>
              <button @click="editBook(book)">Edit</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Edit Book Form -->
    <div v-if="editingBook">
      <h2>Edit Book</h2>
      <form @submit.prevent="updateBook">
        <input v-model="editingBook.name" placeholder="Book Name" required />
        <input v-model="editingBook.author" placeholder="Author" required />
        <input v-model="editingBook.publication" placeholder="Publication" required />
        <button type="submit" :disabled="isUpdating">Update Book</button>
      </form>
    </div>
  </div>
</template>

<script>
import BookService from './services/BookService';

export default {
  data() {
    return {
      books: [],
      newBook: {
        name: '',
        author: '',
        publication: ''
      },
      editingBook: null,
      isCreating: false,
      isUpdating: false
    };
  },

  created() {
    this.fetchBooks();
  },

  methods: {
    async fetchBooks() {
      console.log('Fetching books...');
      try {
        const response = await BookService.getAllBooks();
        this.books = response.data;
        console.log('Fetched books:', this.books);  // Debugging line
      } catch (error) {
        console.error('Error fetching books:', error);
      }
    },

    async createBook() {
      this.isCreating = true;
      try {
        await BookService.createBook(this.newBook);
        console.log('Book created successfully');  // Debugging line
        await this.fetchBooks();  // Ensure fetchBooks is awaited
        this.newBook = { name: '', author: '', publication: '' };
      } catch (error) {
        console.error('Error creating book:', error);
      } finally {
        this.isCreating = false;
      }
    },

    confirmDelete(bookID) {
      // Show confirmation dialog
      const confirmed = window.confirm('Are you sure you want to delete this book?');
      
      if (confirmed) {
        this.deleteBook(bookID);
      }
    },

    async deleteBook(bookID) {
      console.log('Deleting book with ID:', bookID);  // Debugging line
      if (!bookID) {
        console.error('Error: bookID is undefined or null');
        return;
      }
      
      try {
        await BookService.deleteBook(bookID);
        console.log('Book deleted successfully');  // Debugging line
        await this.fetchBooks();  // Refresh the book list after deleting
      } catch (error) {
        console.error('Error deleting book:', error);
      }
    },

    editBook(book) {
      this.editingBook = { ...book };
    },

    async updateBook() {
  this.isUpdating = true;
  try {
    await BookService.updateBook(this.editingBook.ID, this.editingBook);
    console.log('Book updated successfully');
    await this.fetchBooks();  // Refresh the book list after updating
    this.editingBook = null;
  } catch (error) {
    console.error('Error updating book:', error.response?.data || error.message);
    alert('Error updating book. Please try again.');  // Show user-friendly message
  } finally {
    this.isUpdating = false;
  }
}

  }
};
</script>

<style scoped>
#app {
  font-family: Arial, sans-serif;
  margin: 20px;
}

form {
  margin-bottom: 20px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  padding: 10px;
  border: 1px solid #ddd;
}

button {
  margin-right: 10px;
}

button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
</style>
