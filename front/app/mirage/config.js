import Mirage from 'ember-cli-mirage';

export default function () {

  this.post('/api/token-auth', function (db, request) {
    var data = JSON.parse(request.requestBody),
      success = true;

    if (success) {
      return {token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ'};
    } else {
      return new Mirage.Response(401, {}, {});
    }
  });

  this.get('/notices', function () {
    let notices = [
      {
        type: 'notice',
        id: 1,
        attributes: {
          'content': 'Online Reading',
          'created-at': (new Date()).toISOString()
        }
      },
      {
        type: 'notice',
        id: 2,
        attributes: {
          'content': 'Update username',
          'created-at': (new Date(new Date().setDate(new Date().getDate() - 1))).toISOString()
        }
      },
      {
        type: 'notice',
        id: 3,
        attributes: {
          'content': 'Change password',
          'created-at': (new Date(new Date().setDate(new Date().getDate() - 6))).toISOString()
        }
      }
    ];

    return {data: notices};
  });

  this.get('/logs', function () {
    let logs = [
      {
        type: 'log',
        id: 1,
        attributes: {
          'content': 'Sign in',
          'created-at': (new Date()).toISOString()
        }
      },
      {
        type: 'log',
        id: 2,
        attributes: {
          'content': 'Read a book',
          'created-at': (new Date(new Date().setDate(new Date().getDate() - 1))).toISOString()
        }
      },
      {
        type: 'log',
        id: 3,
        attributes: {
          'content': 'Add a new book',
          'created-at': (new Date(new Date().setDate(new Date().getDate() - 6))).toISOString()
        }
      }
    ];

    return {data: logs};
  })
}
