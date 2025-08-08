import React from "react";
import Layout from "../../../layout/Layout";
import PageMeta from "../../common/PageMeta";
import Navbar from "../../../layout/Header/Navbar";
import PageHeader from "../../common/PageHeader";
import BlogGrid from "../../blogs/BlogGrid";
import FooterTwo from "../../../layout/Footer/FooterTwo";

const Blogs = () => {
  return (
    <Layout>
      <PageMeta title="Welcome Our Blog- Software &amp; IT Solutions HTML Templat" />
      <Navbar navDark />
      <PageHeader
        title="Our Latest News and Blogs"
        desc="Completely integrate equity invested partnerships without revolutionary systems. Monotonectally network pandemic e-services via bricks-and-clicks information."
        blogtags
      />
      <BlogGrid />
      <FooterTwo />
    </Layout>
  );
};

export default Blogs;
