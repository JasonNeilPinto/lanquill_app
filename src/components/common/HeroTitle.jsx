/* eslint-disable react/prop-types */
import React from "react";

const HeroTitle = ({ subtitle, title, desc }) => {
  return (
    <>
      {subtitle ? <h5 style={{ color: "#549D54" }}>{subtitle}</h5> : ""}
      <h1 className="fw-bold display-5" data-aos="fade-up">
        {title}
      </h1>
      <p className="lead" data-aos="fade-up" data-aos-delay="50">
        {desc}
      </p>
    </>
  );
};

export default HeroTitle;
