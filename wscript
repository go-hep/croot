#

top = '.'
out = '__build__'

def options(ctx):
    pass

def configure(ctx):
    ctx.load('go')
    
def build(ctx):

    ctx(
        features='cgopackage',
        name ='go-croot',
        source='pkg/croot.go',
        target='croot',
        use = [
            'croot',
            ],
        )
    
