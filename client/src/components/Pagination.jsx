import React from 'react';

import { useLocation, useNavigate } from 'react-router-dom';


const Pagination = ({ numOfPages, page }) => {
  const pagesArray = Array.from({ length: numOfPages }, (_, i) => i + 1);
  // const dispatch = useDispatch();
  const { pathname, search } = useLocation();
  const navigate = useNavigate();

  const handlePageChange = (newPage) => {
    // dispatch(setState({ name: 'page', value: newPage }));
    // navigate('/products');

    // new method
    let searchParams = new URLSearchParams(search);
    searchParams.set("page", newPage);
    // console.log(searchParams);
    navigate(`${pathname}?${searchParams}`);
  };

  const renderPagination = () => {
    if (numOfPages < 5) {
      return pagesArray.map((pageNumber) => (
        <button
          key={pageNumber}
          onClick={() => handlePageChange(pageNumber)}
          className={`px-3 py-1 border rounded mx-1 ${
            page === pageNumber ? 'bg-blue-500 text-white' : 'bg-white'
          }`}
        >
          {pageNumber}
        </button>
      ));
    }

    if (page === 1 || page === numOfPages) {
      return (
        <>
          <button
            className="px-3 py-1 border rounded mx-1"
            onClick={() => handlePageChange(1)}
          >
            {1}
          </button>
          {page === 1 ? (
            <>
              <button
                className="px-3 py-1 border rounded mx-1"
                onClick={() => handlePageChange(2)}
              >
                {2}
              </button>
              <span className="px-3 py-1">...</span>
              <button
                className="px-3 py-1 border rounded mx-1"
                onClick={() => handlePageChange(numOfPages)}
              >
                {numOfPages}
              </button>
            </>
          ) : (
            <>
              <span className="px-3 py-1">...</span>
              <button
                className="px-3 py-1 border rounded mx-1"
                onClick={() => handlePageChange(numOfPages)}
              >
                {numOfPages}
              </button>
            </>
          )}
        </>
      );
    }

    if (page === 2) {
      return (
        <>
          <button
            className="px-3 py-1 border rounded mx-1"
            onClick={() => handlePageChange(1)}
          >
            {1}
          </button>
          <button
            className="px-3 py-1 border rounded mx-1"
            onClick={() => handlePageChange(2)}
          >
            {2}
          </button>
          <button
            className="px-3 py-1 border rounded mx-1"
            onClick={() => handlePageChange(3)}
          >
            {3}
          </button>
          <span className="px-3 py-1">...</span>
          <button
            className="px-3 py-1 border rounded mx-1"
            onClick={() => handlePageChange(numOfPages)}
          >
            {numOfPages}
          </button>
        </>
      );
    }

    return (
      <>
        <button
          className="px-3 py-1 border rounded mx-1"
          onClick={() => handlePageChange(1)}
        >
          {1}
        </button>
        <span className="px-3 py-1">...</span>
        <button
          className="px-3 py-1 border rounded mx-1"
          onClick={() => handlePageChange(page - 1)}
        >
          {page - 1}
        </button>
        <button
          className="px-3 py-1 border rounded mx-1"
          onClick={() => handlePageChange(page)}
        >
          {page}
        </button>
        <button
          className="px-3 py-1 border rounded mx-1"
          onClick={() => handlePageChange(page + 1)}
        >
          {page + 1}
        </button>
        <span className="px-3 py-1">...</span>
        <button
          class="px-3 py-1 border rounded mx-1"
          onClick={() => handlePageChange(numOfPages)}
        >
          {numOfPages}
        </button>
      </>
    );
  };

  return (
    <div className="flex items-center justify-end mt-6">
      <button
        className="px-3 py-1 border rounded mx-1"
        onClick={() => {
          if (page === 1) {
            handlePageChange(numOfPages);
          } else {
            handlePageChange(page - 1);
          }
        }}
      >
        Prev
      </button>
      {renderPagination()}
      <button
        className="px-3 py-1 border rounded mx-1"
        onClick={() => {
          if (page === numOfPages) {
            handlePageChange(1);
          } else {
            handlePageChange(page + 1);
          }
        }}
      >
        Next
      </button>
    </div>
  );
};

export default Pagination;
