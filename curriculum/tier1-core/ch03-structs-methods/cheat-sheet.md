# Chapter 03 Cheat Sheet: Structs, Methods & Receivers

| Pattern | Code |
|---------|------|
| Define a struct | `type User struct { Name string; Age int }` |
| Create with named fields | `u := User{Name: "Alice", Age: 30}` |
| Access a field | `u.Name` |
| Value receiver method | `func (u User) Greet() string { return "Hi, " + u.Name }` |
| Pointer receiver method | `func (u *User) SetName(n string) { u.Name = n }` |
| Embed a struct | `type Admin struct { User; Role string }` |
| Access promoted field | `admin.Name` (promoted from `User`) |
| JSON struct tag | `` `json:"user_name"` `` |
| Omit zero-value field in JSON | `` `json:"age,omitempty"` `` |
| Exclude field from JSON | `` `json:"-"` `` |
| Zero-value struct | `var u User` → `Name: "", Age: 0` |
| Constructor function | `func NewUser(name string) (*User, error) { ... }` |
| Print struct with field names | `fmt.Printf("%+v\n", u)` |
