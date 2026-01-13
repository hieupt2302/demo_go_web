export interface Category {
  id: number;
  name: string;
  description?: string;
  books?: Book[];
}

export interface Author {
  id: number;
  name: string;
  bio?: string;
  books?: Book[];
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
  category_ids: number[]; 
  author_ids: number[];
  categories?: Category[];
  authors?: Author[];
}