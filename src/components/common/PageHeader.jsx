/* eslint-disable react/prop-types */
import React from "react";
// import { Link } from "react-router-dom";

const PageHeader = ({ title, desc }) => {
  return (
    <>
      <section
        className="page-header position-relative overflow-hidden ptb-120 bg-dark"
        style={{
          background: "url('/img/page-header-bg.svg')no-repeat bottom left",
        }}
      >
        <div className="container">
          <div className="col-lg-8 col-md-12">
            <h1 className="display-5 fw-bold">{title}</h1>
            <p className="lead">{desc}</p>
          </div>
        </div>

        <div className="bg-circle rounded-circle circle-shape-3 position-absolute bg-dark-light right-5"></div>
      </section>
    </>
  );
};

export default PageHeader;
