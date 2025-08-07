import React from "react";
import PageMeta from "../../common/PageMeta";
import ContactBox from "../../contact/ContactBox";
import ContactFormTwo from "../../contact/ContactFormTwo";
import FooterOne from "../../../layout/Footer/FooterOne";
import PageHeader from "../../common/PageHeader";
import Navbar from "../../../layout/Header/Navbar";
import Layout from "../../../layout/Layout";

const Contact = () => {
  return (
    <Layout>
      <PageMeta title="Contact-Software &amp; IT Solutions HTML Template" />
      <Navbar classOption="navbar-light" />
      <PageHeader
        title="Contact Us"
        desc="Seamlessly actualize client-based users after out-of-the-box value data through frictionless expertise. Proactively coordinate quality quality vectors vis-a-vis supply chains. Quickly engage client-centric web services."
      />
      <ContactBox />
      <ContactFormTwo />
      <FooterOne
        style={{
          background: "url('/img/page-header-bg.svg')no-repeat bottom right",
        }}
      />
    </Layout>
  );
};

export default Contact;
