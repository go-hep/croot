#

top = '.'
out = '__build__'

def options(ctx):
    pass

def configure(ctx):
    ctx.load('go')
    
def build(ctx):

    o=ctx(
        features='cgopackage',
        name ='go-croot',
        source='''
        pkg/croot/croot.go
        pkg/croot/croot_genreflex.go
        pkg/croot/croot_reflex.go
        pkg/croot/croot_cintex.go
        pkg/croot/croot_cstruct.go
        ''',
        target='bitbucket.org/binet/go-croot/pkg/croot',
        use = [
            'croot',
            'go-ctypes',
            ],
        )
    #o.env.GOMAKE_FLAGS=' '
    
    ctx(
        features='go goprogram',
        name   = 'test-croot-ex-tree-00',
        source ='cmd/test-croot-ex-tree-00.go',
        target = 'test-croot-ex-tree-00',
        use = ['go-croot',],
        )
