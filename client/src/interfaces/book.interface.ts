export interface Category {
  id: number;
  name: string;
  description?: string;
}

export interface Author {
  id: number;
  name: string;
  bio?: string;
}

export interface Book {
  id: number;
  title: string;
  isbn: string;
  price: number;
  stock_quantity: number;
  publisher?: string;
  publish_date?: string;
  cover_image_url: string;
  description: string;
  category_id: number;
  author_id: number;
  category?: Category;
  author?: Author;
}