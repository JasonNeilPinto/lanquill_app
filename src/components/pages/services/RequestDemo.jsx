import React from "react";
import Layout from "../../../layout/Layout";
import PageMeta from "../../common/PageMeta";
import Navbar from "../../../layout/Header/Navbar";
import ContactsForm from "../../contact/ContactsForm";
import FooterTwo from "../../../layout/Footer/FooterTwo";
import Review from "../../others/ReviewTab";

const RequestDemo = () => {
  return (
    <Layout>
      <PageMeta title="Request for Demo - Software &amp; IT Solutions HTML Template" />
      <Navbar navDark posAbsolute />
      <section
        className="sign-up-in-section bg-dark ptb-120"
        style={{
          background: "url('/img/page-header-bg.svg')no-repeat bottom right",
        }}
      >
        <div className="container">
          <div className="row align-items-center justify-content-between">
            <ContactsForm />
            <Review reqPage />
          </div>
        </div>
      </section>
      <FooterTwo />
    </Layout>
  );
};

export default RequestDemo;
