```
nvm use v19.8.1
npx create-next-app@latest .

npm run dev
```

Youtube playlist - https://www.youtube.com/watch?v=ZjAqacIC_3c&list=PLC3y8-rFHvwjOKd6gdf4QtV1uYNiQnruI

NextJS is a React framework to build web apps

**React**

- not feasible to production ready web apps
- React is a library to build user interfaces
- We need to pick up different tools for routing, data fetching etc.

**Next.js**

- to build production ready apps
- features including routing, data fetching, bundling and compiling
- opinions and conventions should be followed

- simplifies the process of building a web application for production

1. Routing
2. API routes
3. Rendering - supports both server side rendering and client side rendering
4. Simplified data fetching
5. Styling
6. Optimization for images, fonts, scripts
7. Dev and prod build system

**Flow of control**

- layout.tsx -> RootLayout component is rendered
- localhost:3000 -> children prop will always refer to the component defined in page.tsx

**React Server Components (RSC)**

- React server components is a new architecture introduced in version 18 embraced by NextJS
- RSC introduces a new way of creating React components, splitting them into two types:
  - server components
    - all components by default are server components
    - server components can run tasks like reading files, fetching data from database
    - cannot use hooks or handle user interactions
  - client components
    - add `use client` at the top of the component file to create a client component
    - can't perform tasks like reading files and reading data form db
    - can use hooks and manage interactions
    - traditional react components

### Routing

- NextJS has a file-system based routing mechanism
- URL paths that users can access in the browser are defined by files and folders in your codebase

**Routing conventions**

- All routes must be placed inside the app folder
- Every file that corresponds to a route must be named `page.js` or `page.tsx`
- Every folder corresponds to a pth segment in the browser URL

**Nested Routes**

- Achieved using nested folders

**Dynamic Routes**

- Products listing and detailed page

```
/products
/products/1

products/[id]
```

Every page receives route parameter as a prop

```
export default function ProductDetails({params}) {
  return (
    <>
    <h1>Product List {params.productId}</h1>
    </>
  )
}
```

**Nested Dynamic Routes**

```
/products/1
/products/1/reviews/1

reviews folder in [productID]
```

![folder structure](nested_dynamic_routes.png)

**Catch-all Segments**

- captures all url segments and maps them to single file `page.tsx`

**Custom 404 page**

- create `not-found.tsx` in `app` directory
- `notFound()` function
- `import {notFound} from "next/navigation";`
- we can create `not-found.tsx` at route specific folders

### File Colocation

- until we have a `page.tsx` file in the folder, it is not publicly accessible
- only the content in `page.tsx` is sent to the client

**Private folders**

- A private folder indicates that it is a private implementation detail and should not be considered by the routing system
- The folders and all its subfolders are excluded from routing
- Prefix the folder name with an underscore
- Useful
  - For separating UI logic from routing logic
  - For consistently organizing internal files across a project
- If you want to include an underscore in URL segments, you can prefix the folder name with `%5F` which is the URL encoded form of an underscore.

### Route Groups

- Allows us to logically group our routes and project files without affecting the URL path structure
- Authentication routes
  - Register
  - Login
  - Forgot Password
- All routes are put in `auth` folder in `app` folder.
- Wrap `auth` in paranthesis `()`. It looks like `(auth)`. NextJS will omit `auth` from the URL.

### Layouts

- A page is UI that is unique to a route
- A layout is UI that is shared between multiple pages in the app

- Define a layout by default exporting a React component from a `layout.js` or `layout.tsx` file
- The component should accept a children prop that will be populated with a child page during rendering
- Root layout is at `app` directory. NextJS will recreate even if we delete the file

**Nested layouts**

- `layout.tsx` in specific route directory

**Route group layout**

- To selectively apply a layout to certain segments while leaving others unchanged

### Routing Metadata

- SEO
- Next.js introduces the Metadata API which allows you to define metadata for each page
- Metadata ensures accurate and relevant information is displayed when your pages are shared and indexed
- we can configure metadata in `page.tsx` or `layout.tsx`

- Export a static metadata object
- Export a dynamic generateMetadata function

- Both `layout.tsx` and `page.tsx` files can export metadata. Layout metadata is applied to all pages.
  Page metadata is applied only to the pages
- Metadata is read in order, from the root level down to the final page level
- When there is metadata in multiple places for the same route, they get combined, but page metadata will
  replace layout metadata if they have the same properties

**Static metadata object**

```js
export const metadata = {
  title: "Learning NextJS",
  description: "Generated by me",
};
```

**Dynamic metadata**

- Common usecase for dynamic metadata is usually for dynamic routes
  eg: products

  ```js
  import {Metadata} from "next";

  type Props = {
  params:{
  productId: string;
  }
  }

  export const generateMetadata = ({params}: Props): Metadata => {
  return {
  title: `Product ${params.productId}`
  }
  ```

- geneateMetadata can be async

```js
export const generateMetadata = async ({
  params,
}: Props): Promise<Metadata> => {
  const title = await new Promise((resolve) => {
    setTimeout(() => {
      resolve(`iPhone ${params.productId}`);
    }, 100);
  });
  return {
    title: `Product ${title}`,
  };
};
```

start from - title metadata - https://www.youtube.com/watch?v=1OqftoKO2V0&list=TLPQMTUwNTIwMjQrlBvc3jMeHw&index=18
