import React from "react";
import Layout from "../../../layout/Layout";
import PageMeta from "../../common/PageMeta";
import Navbar from "../../../layout/Header/Navbar";
import PageHeader from "../../common/PageHeader";
import LatestBlog from "../../blogs/LatestBlog";
import CtaTwo from "../../cta/CtaTwo";
import Ecommerce from "../../blogs/Ecommerce";
import FooterTwo from "../../../layout/Footer/FooterTwo";

const EcommerceBlog = () => {
  return (
    <Layout>
      <PageMeta title="Blog Details - Software &amp; IT Solutions HTML Template" />
      <Navbar navDark />
      <PageHeader title="Cybersecurity Best Practices for Protecting Customer Data in E-commerce: A Guide for CTOs" />
      <Ecommerce />
      {/* <NewsLetter /> */}
      <LatestBlog />
      <CtaTwo />
      <FooterTwo />
    </Layout>
  );
};

export default EcommerceBlog;
