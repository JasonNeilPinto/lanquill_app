import React from "react";
import { Helmet } from "react-helmet";
import PropTypes from "prop-types";

const PageMeta = ({ title }) => {
  return (
    <Helmet>
      <title>{title}</title>
    </Helmet>
  );
};

PageMeta.propTypes = {
  title: PropTypes.string.isRequired,
};

export default PageMeta;
