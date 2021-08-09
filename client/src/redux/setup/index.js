if (process.env.NODE_ENV === 'development'){
  module.exports = require('./configStore.dev')
} else {
    module.exports = require('./configStore.prod')
}



