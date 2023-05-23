// Write your MyBook class here
class MyBook: public Book {
    private:
        int price;
    public:
        //   Class Constructor
        //   
        //   Parameters:
        //   title - The book's title.
        //   author - The book's author.
        //   price - The book's price.
        //
        // Write your constructor here
        MyBook(string title, string author, int price): Book(title, author) {
            this->title = title;
            this->author = author;
            this->price = price;
        }
    
        //   Function Name: display
        //   Print the title, author, and price in the specified format.
        //
        // Write your method here
        void display() {
            cout << "Title: " << this->title << "\n";
            cout << "Author: " << this->author << "\n";
            cout << "Price: " << this->price << "\n";
        }
};
// End class