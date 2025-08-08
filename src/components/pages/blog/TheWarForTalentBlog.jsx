import React from "react";
import Layout from "../../../layout/Layout";
import PageMeta from "../../common/PageMeta";
import Navbar from "../../../layout/Header/Navbar";
import PageHeader from "../../common/PageHeader";
import LatestBlog from "../../blogs/LatestBlog";
import CtaTwo from "../../cta/CtaTwo";
import TheWarForTalent from "../../blogs/TheWarForTalent";
import FooterTwo from "../../../layout/Footer/FooterTwo";

const TheWarForTalentBlog = () => {
  return (
    <Layout>
      <PageMeta title="Blog Details - Software &amp; IT Solutions HTML Template" />
      <Navbar navDark />
      <PageHeader title="The War for Talent: How IT Staffing Can Be Your Secret Weapon" />
      <TheWarForTalent />
      {/* <NewsLetter /> */}
      <LatestBlog />
      <CtaTwo />
      <FooterTwo />
    </Layout>
  );
};

export default TheWarForTalentBlog;
