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
}
